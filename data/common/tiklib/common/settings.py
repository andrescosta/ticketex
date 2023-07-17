from typing import Optional
from pydantic import BaseSettings


class Settings(BaseSettings):
    # database configurations
    DATABASE_URL: Optional[str] = None
    DATABASE_NAME: Optional[str] = None
    AUTH_PROV_DOMAIN: Optional[str] = None
    AUTH_PROV_AUD: Optional[str] = None
    
    class Config:
        env_file = '.env'
        env_file_encoding = 'utf-8'
        auth_prov_domain = ''
        auth_prov_aud = ''

app_settings = Settings()