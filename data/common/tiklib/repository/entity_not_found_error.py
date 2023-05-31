from typing import Optional


class EntityNotFoundError(Exception):
    entity_type:type

    def __init__(self, id:str,entity_type:type ):
        self.id = id
        self.entity_type = entity_type
    
    def __str__(self) -> str:
        return f'{self.entity_type.__name__} not found.'   
