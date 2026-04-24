from sqlalchemy import func, select
from sqlalchemy.orm import Session

from app import models


STATUS_ORDER = ["pending", "processing", "completed"]


def generate_study_plan(learner_name: str, topic: str, level: str, duration_days: int, goals: str) -> str:
    daily_focus = max(1, duration_days // 3)
    return (
        f"Plano para {learner_name}:\n"
        f"1. Revisar os fundamentos de {topic} por {daily_focus} dias, criando resumos curtos.\n"
        f"2. Separar blocos praticos para o nivel {level} e medir progresso por entregas semanais.\n"
        f"3. Transformar o objetivo '{goals}' em uma checklist com marcos diarios.\n"
        f"4. Reservar o ultimo terco do cronograma para revisao, correcao de lacunas e demonstracao final do aprendizado."
    )


def build_stats(db: Session) -> dict[str, int]:
    total_requests = db.scalar(select(func.count()).select_from(models.StudyPlanRequest)) or 0
    pending_requests = db.scalar(
        select(func.count()).select_from(models.StudyPlanRequest).where(models.StudyPlanRequest.status == "pending")
    ) or 0
    processing_requests = db.scalar(
        select(func.count()).select_from(models.StudyPlanRequest).where(models.StudyPlanRequest.status == "processing")
    ) or 0
    completed_requests = db.scalar(
        select(func.count()).select_from(models.StudyPlanRequest).where(models.StudyPlanRequest.status == "completed")
    ) or 0

    return {
        "total_requests": total_requests,
        "pending_requests": pending_requests,
        "processing_requests": processing_requests,
        "completed_requests": completed_requests,
    }
