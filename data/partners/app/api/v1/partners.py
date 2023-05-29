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
    partner = await repository.partner.get(Partner(id=partner_id))
    return partner


@router.put("/{partner_id}", response_model=Partner)
async def put(partner_id: str, upartner: Partner):
    upartner.id = partner_id
    newp = await repository.partner.update(upartner)
    return newp


@router.delete("/{partner_id}")
async def delete(partner_id: str):
    await repository.partner.delete(Partner(id=partner_id))
    return Response()
