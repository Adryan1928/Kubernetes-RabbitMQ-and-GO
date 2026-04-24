from datetime import datetime

from pydantic import BaseModel, Field


class StudyPlanCreate(BaseModel):
    learner_name: str = Field(min_length=2, max_length=120)
    topic: str = Field(min_length=3, max_length=200)
    level: str = Field(min_length=2, max_length=40)
    duration_days: int = Field(ge=1, le=90)
    goals: str = Field(min_length=10, max_length=1000)


class StudyPlanResponse(BaseModel):
    id: int
    learner_name: str
    topic: str
    level: str
    duration_days: int
    goals: str
    status: str
    generated_plan: str | None
    created_at: datetime
    updated_at: datetime

    model_config = {"from_attributes": True}


class StatsResponse(BaseModel):
    total_requests: int
    pending_requests: int
    processing_requests: int
    completed_requests: int
