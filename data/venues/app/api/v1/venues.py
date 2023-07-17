import logging
from fastapi import APIRouter
from fastapi.responses import Response
from app import repository
from app.models import Venue

router = APIRouter()

@router.post("", response_model=Venue)
async def post(*, venue: Venue):
    newvenue = await repository.venue.save(venue)
    return newvenue


@router.get("/{venue_id}", response_model=Venue)
async def get(venue_id: str):
    res = await repository.venue.get(venue_id)
    return res

@router.get("")
async def get():
    res = await repository.venue.get(None)
    return res

@router.put("/{venue_id}")
async def put(venue_id: str, venue: Venue):
    res = await repository.venue.update(venue_id, venue)
    return res


@router.delete("/{venue_id}")
async def delete(venue_id: str):
    await repository.venue.delete(venue_id)
    return Response()
