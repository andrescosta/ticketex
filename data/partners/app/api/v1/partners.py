from fastapi import APIRouter
from fastapi.responses import Response
from app import repository
from app.models import Partner

router = APIRouter()


@router.post("", response_model=Partner)
async def post(*, partner: Partner):
    newpartner = await repository.partner.save(partner)
    return newpartner


@router.get("/{partner_id}", response_model=Partner)
async def get(partner_id: str):
    res = await repository.partner.get(partner_id)
    return res


@router.put("/{partner_id}")
async def put(partner_id: str, partner: Partner):
    res = await repository.partner.update(partner_id, partner)
    return res


@router.delete("/{partner_id}")
async def delete(partner_id: str):
    await repository.partner.delete(partner_id)
    return Response()
