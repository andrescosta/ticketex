import logging
from fastapi import APIRouter, Depends
from fastapi.responses import Response
from app import repository
from app.models import Venue
from tiklib.api.auth.jwt_bearer import JWTBearer, JWTClaimsCredential
from typing import Annotated

router = APIRouter()

@router.post("", response_model=Venue)
async def post(*, venue: Venue, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    venue.partner_id = token.sub
    newvenue = await repository.venue.save(venue)
    return newvenue


@router.get("/{venue_id}", response_model=Venue)
async def get(venue_id: str, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    res = await repository.venue.get(token.sub, venue_id)
    return res

@router.get("")
async def get(token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    res = await repository.venue.get(token.sub, None)
    return res

@router.put("/{venue_id}")
async def put(venue_id: str, venue: Venue, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    venue.partner_id = token.sub
    res = await repository.venue.update(venue_id, venue)
    return res


@router.delete("/{venue_id}")
async def delete(venue_id: str, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    await repository.venue.delete(token.sub, venue_id)
    return Response()
