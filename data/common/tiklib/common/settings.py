from typing import Optional
from pydantic import BaseSettings


class Settings(BaseSettings):
    # database configurations
    DATABASE_URL: Optional[str] = None
    DATABASE_NAME: Optional[str] = None
    
    class Config:
        env_file = '.env'
        env_file_encoding = 'utf-8'

app_settings = Settings()