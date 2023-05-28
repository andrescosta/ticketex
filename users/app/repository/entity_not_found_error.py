class EntityNotFoundError(Exception):
    entity_type = None
    def __init__(self, id,entity_type ):
        self.id = id
        self.entity_type = entity_type
    
    def __str__(self) -> str:
        return f'{self.entity_type.__name__} not found.'   
