from tiklib.api.auth.auth_exceptions import PermissionDeniedException
from tiklib.common.settings import app_settings
import jwt
import logging
from typing import Any

logger = logging.getLogger(__name__)

class JWTDecoder:
    _instance=None 
    auth0_issuer_url: str = f"https://{app_settings.AUTH_PROV_DOMAIN}/"
    auth0_audience: str = app_settings.AUTH_PROV_AUD
    algorithm: str = "RS256"
    jwks_uri: str = f"{auth0_issuer_url}.well-known/jwks.json"
    jwks_client:jwt.PyJWKClient

    def decode(self, jwtoken: str )->Any:
        jwt_signing_key = self.jwks_client.get_signing_key_from_jwt(
                jwtoken
        ).key
        decoded = jwt.decode(
            jwtoken,
            jwt_signing_key,
            algorithms=self.algorithm,
            audience=self.auth0_audience,
            issuer=self.auth0_issuer_url,
        )
        return decoded
    
    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(JWTDecoder, cls).__new__(cls)
            cls._instance.jwks_client = jwt.PyJWKClient(cls._instance.jwks_uri)
        return cls._instance

