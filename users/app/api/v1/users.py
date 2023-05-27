from fastapi import APIRouter, HTTPException
from fastapi.responses import Response
from bson import json_util
from app import models, repository
from app.repository.entity_not_found_error import EntityNotFoundError
from app.models import User
router = APIRouter()

@router.post("/", response_model=User)
async def post(*,user: User):
    newuser = await repository.user.save(user)
    return newuser

@router.get("/{user_id}")
async def get(user_id: str):
    try:
        user = await repository.user.get(user_id)
        return user
    except EntityNotFoundError as e:
        raise HTTPException(status_code=404, detail=str(e))

@router.put("/{user_id}")
async def put(user_id: str, user: User):
    try:
        user.id = user_id
        await repository.user.update(user)
        return user
    except EntityNotFoundError as e:
        raise HTTPException(status_code=404, detail=str(e))

@router.delete("/{user_id}")
async def delete(user_id: str):
    try:
        await repository.user.delete(user_id)
        return Response()
    except EntityNotFoundError as e:
        raise HTTPException(status_code=404, detail=str(e))
