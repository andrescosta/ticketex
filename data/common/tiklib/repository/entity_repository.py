from typing import Any
from pymongo.collection import Collection
import copy
from bson import ObjectId
from tiklib.db.client import get_db



from tiklib.models.entity import Entity
from tiklib.repository.entity_not_found_error import EntityNotFoundError


class EntityRepository:
 
    def __init__(self,entity_type:type) -> None:
        self.entity_type=entity_type
        
    async def save(self, entity:Entity)->Entity:
        entity_dict = self.for_saving(entity)
        result = await self.collection().insert_one(entity_dict)
        nentity = copy.copy(entity)
        nentity.id = str(result.inserted_id); 
        return nentity
    
    async def get(self, id:str):
        entity_dic = await self.collection().find_one({"_id": ObjectId(id)})
        if entity_dic:
            entity_dic["id"] = str(entity_dic.pop("_id"))
            return self.entity_type(**entity_dic)
        else:
            raise EntityNotFoundError(id, self.entity_type)

    async def update(self, id:str, entity:Entity)->Entity:
        entity_dict = self.for_saving(entity)
        result = await self.collection().update_one({"_id": ObjectId(id)}, {"$set": entity_dict})
        if result.matched_count==0:
            raise EntityNotFoundError(id, self.entity_type)
        else:
            return await self.get(id)

    async def delete(self, id:str)->None:
        result = await self.collection().delete_one({"_id": ObjectId(id)})
        if result.deleted_count == 0:
            raise EntityNotFoundError(id, self.entity_type)
    
    def collection_name(self, entitytype:type)->str:
        return f"{entitytype.__name__.lower()}s"
    
    def collection(self)->Collection:
        return get_db()[self.collection_name(self.entity_type)]

    def for_saving(self, entity:Entity)->dict[str, Any]:
        entity_dict = entity.dict(exclude_unset=True)
        entity_dict.pop("id", None)
        return entity_dict