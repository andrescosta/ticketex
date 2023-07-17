from fastapi import APIRouter, Depends
from fastapi.responses import Response
from app import repository
from app.models import Adventure
from tiklib.api.auth.jwt_bearer import JWTBearer, JWTClaimsCredential
from typing import Annotated

router = APIRouter()

@router.post("", response_model=Adventure)
async def post(adventure: Adventure, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    adventure.partner_id = token.sub
    res = await repository.adventure.save(adventure)
    return res

@router.get("/{adventure_id}")
async def get(adventure_id: str, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    partner_id = token.sub
    res = await repository.adventure.get(partner_id, adventure_id)
    return res

@router.get("")
async def get(token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    partner_id = token.sub
    res = await repository.adventure.get(partner_id, None)
    return res

@router.put("/{adventure_id}")
async def put(adventure_id: str, adventure: Adventure, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    adventure.partner_id = token.sub
    res = await repository.adventure.update(adventure_id, adventure)
    return res

@router.delete("/{adventure_id}")
async def delete(adventure_id: str, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    partner_id = token.sub
    await repository.adventure.delete(partner_id, adventure_id)
    return Response()
