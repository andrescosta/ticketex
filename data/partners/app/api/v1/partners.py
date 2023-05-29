from fastapi import APIRouter
from fastapi.responses import Response
from app import repository
from app.models import Partner

router = APIRouter()


@router.post("", response_model=Partner)
async def post(*, partner: Partner):
    print("aaa")
    newpartner = await repository.partner.save(partner)
    return newpartner


@router.get("/{partner_id}")
async def get(partner_id: str):
    partner = await repository.partner.get(partner_id)
    return partner


@router.put("/{partner_id}")
async def put(partner_id: str, partner: Partner):
    await repository.partner.update(partner_id, partner)
    return partner


@router.delete("/{partner_id}")
async def delete(partner_id: str):
    await repository.partner.delete(partner_id)
    return Response()
