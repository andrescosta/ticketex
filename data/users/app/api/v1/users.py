from fastapi import APIRouter
from fastapi.responses import Response
from app import repository
from app.models import User

router = APIRouter()

@router.post("", response_model=User)
async def post(*, user: User):
    newuser = await repository.user.save(user)
    return newuser


@router.get("/{user_id}")
async def get(user_id: str):
    res = await repository.user.get(user_id)
    return res


@router.put("/{user_id}")
async def put(user_id: str, user: User):
    res = await repository.user.update(user_id, user)
    return res


@router.delete("/{user_id}")
async def delete(user_id: str):
    await repository.user.delete(user_id)
    return Response()
