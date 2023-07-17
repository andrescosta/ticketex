from fastapi import Request
from fastapi.security import HTTPBearer, HTTPAuthorizationCredentials
from tiklib.api.auth.auth_exceptions import PermissionDeniedException
from tiklib.api.auth.jwt_validator import JWTDecoder
import logging
from pydantic import BaseModel
from typing import Any


logger = logging.getLogger(__name__)

jwt_decoder = JWTDecoder()

class JWTClaimsCredential:
    sub:str
    claims:Any
    def __init__(self, claims):
        for key in claims:
            setattr(self, key, claims[key])
        self.claims = claims 

class JWTBearer(HTTPBearer):

    def __init__(self, auto_error: bool = True):
        super(JWTBearer, self).__init__(auto_error=auto_error)

    async def __call__(self, request: Request):
        credentials: HTTPAuthorizationCredentials = await super(JWTBearer, self).__call__(request)
        if credentials:
            if not credentials.scheme == "Bearer":
                logger.error("Authorization token is not Bearer")
                raise PermissionDeniedException()
            try:
                claims =  jwt_decoder.decode(credentials.credentials);
                return JWTClaimsCredential(claims)
            except Exception as e:
                logger.debug(e)
                raise PermissionDeniedException from e
        else:
            logger.error("Authorization token empty")
            raise PermissionDeniedException()