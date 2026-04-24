import json

from fastapi import Depends, FastAPI
from fastapi.middleware.cors import CORSMiddleware
from sqlalchemy import select
from sqlalchemy.orm import Session

from app.database import Base, engine, get_db
from app.models import StudyPlanRequest
from app.queue import enqueue_plan, get_redis_client
from app.schemas import StatsResponse, StudyPlanCreate, StudyPlanResponse
from app.services import build_stats
from app.settings import settings

Base.metadata.create_all(bind=engine)

app = FastAPI(
    title="StudyFlow API",
    version="1.0.0",
    description="API que recebe pedidos de planos de estudo e delega o processamento para um worker assincrono.",
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/health")
def healthcheck() -> dict[str, str]:
    return {"status": "ok", "environment": settings.app_env}


@app.get("/jobs", response_model=list[StudyPlanResponse])
def list_jobs(db: Session = Depends(get_db)) -> list[StudyPlanRequest]:
    result = db.scalars(select(StudyPlanRequest).order_by(StudyPlanRequest.created_at.desc()))
    return list(result)


@app.post("/jobs", response_model=StudyPlanResponse, status_code=201)
def create_job(payload: StudyPlanCreate, db: Session = Depends(get_db)) -> StudyPlanRequest:
    job = StudyPlanRequest(**payload.model_dump(), status="pending")
    db.add(job)
    db.commit()
    db.refresh(job)

    enqueue_plan(job.id)
    return job


@app.get("/stats", response_model=StatsResponse)
def get_stats(db: Session = Depends(get_db)) -> dict[str, int]:
    redis_client = get_redis_client()
    cached_stats = redis_client.get(settings.stats_cache_key)
    if cached_stats:
        return json.loads(cached_stats)

    stats = build_stats(db)
    redis_client.setex(settings.stats_cache_key, 15, json.dumps(stats))
    return stats
