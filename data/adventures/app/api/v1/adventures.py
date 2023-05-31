from fastapi import APIRouter
from fastapi.responses import Response
from app import repository
from app.models import Adventure

router = APIRouter()

@router.post("", response_model=Adventure)
async def post(adventure: Adventure):
    res = await repository.adventure.save(adventure)
    return res

@router.get("/{adventure_id}")
async def get(adventure_id: str):
    res = await repository.adventure.get(adventure_id)
    return res


@router.put("/{adventure_id}")
async def put(adventure_id: str, adventure: Adventure):
    res = await repository.adventure.update(adventure_id, adventure)
    return res


@router.delete("/{adventure_id}")
async def delete(adventure_id: str):
    await repository.adventure.delete(adventure_id)
    return Response()
