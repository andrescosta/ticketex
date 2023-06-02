from typing import Any
from pymongo.collection import Collection
import copy
from bson import ObjectId
from tiklib.db.client import get_db
import logging

from tiklib.models.entity import Entity
from tiklib.repository.entity_not_found_error import EntityNotFoundError


logger = logging.getLogger(__name__) 

class EntityRepository:
 
    def __init__(self,entity_type:type) -> None:
        self.entity_type=entity_type
        
    async def save(self, entity:Entity)->Entity:
        entity_dict = self.for_saving(entity)
        result = await self.collection().insert_one(entity_dict)
        nentity = copy.copy(entity)
        nentity.id = str(result.inserted_id); 
        logger.debug(f"The {self.entity_name()} with id {nentity.id} was saved.")
        return nentity
    
    async def get(self, id:str):
        entity_dic = await self.collection().find_one({"_id": ObjectId(id)})
        if entity_dic:
            entity_dic["id"] = str(entity_dic.pop("_id"))
            return self.entity_type(**entity_dic)
        else:
            logger.debug(f"The {self.entity_name()} with id {id} was not found.")
            raise EntityNotFoundError(id, self.entity_type)

    async def update(self, id:str, entity:Entity)->Entity:
        entity_dict = self.for_saving(entity)
        result = await self.collection().update_one({"_id": ObjectId(id)}, {"$set": entity_dict})
        logger.debug(f"The {self.entity_name()} with id {id} was updated.")
        if result.matched_count==0:
            logger.debug(f"The {self.entity_name()} with id {id} was not found.")
            raise EntityNotFoundError(id, self.entity_type)
        else:
            return await self.get(id)

    async def delete(self, id:str)->None:
        result = await self.collection().delete_one({"_id": ObjectId(id)})
        if result.deleted_count == 0:
            logger.debug(f"The {self.entity_name()} with id {id} was not deleted.Not found.")
            raise EntityNotFoundError(id, self.entity_type)
        else:
            logger.debug(f"The {self.entity_name()} with id {id} was deleted.")
    
    def entity_name(self, entitytype:type = None, plural:bool=False)->str:
        if (not entitytype):
            entitytype = self.entity_type
        return f"{entitytype.__name__.lower()}{'s' if plural else ''}"
    
    def collection(self)->Collection:
        return get_db()[self.entity_name(self.entity_type, True)]

    def for_saving(self, entity:Entity)->dict[str, Any]:
        entity_dict = entity.dict(exclude_unset=True)
        entity_dict.pop("id", None)
        return entity_dict
