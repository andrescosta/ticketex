from fastapi import APIRouter, Depends
from fastapi.responses import Response
from app import repository
from app.models import User
import logging
from tiklib.api.auth.jwt_bearer import JWTBearer, JWTClaimsCredential
from typing import Annotated

router = APIRouter()
logger = logging.getLogger(__name__)

@router.post("", response_model=User)
async def post(*, user: User, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    newuser = await repository.user.save(user)
    return newuser

@router.get("")
async def get(token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    res = await repository.user.get(None)
    return res

@router.get("/{user_id}")
async def get(user_id: str, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    res = await repository.user.get(user_id)
    return res

@router.put("/{user_id}")
async def put(user_id: str, user: User, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    uuser = await repository.user.update(user_id, user)
    return uuser


@router.delete("/{user_id}")
async def delete(user_id: str, token: Annotated[JWTClaimsCredential, Depends(JWTBearer())]):
    await repository.user.delete(user_id)
    return Response()
