from pymongo.mongo_client import MongoClient
from pymongo.server_api import ServerApi
from pymongo.mongo_client import MongoClient
from pymongo.server_api import ServerApi
from bson.json_util import dumps
from Movies_pb2 import Movie
import pymongo

class Database:
    def __init__(self):
        uri = ""
        self.client = MongoClient(uri, server_api=ServerApi('1'))
        self.database = self.client.get_database("sample_mflix")
        self.collections = self.database.get_collection("movies")
    
    def format(self, data):
        return dumps(data, indent=4)
        
    def insert(self, movie):
        print(f"[Database] Inserting movie with title {movie.title}")
        return self.collections.insert_one({
            "title": movie.title,
            "directors": movie.directors,
            "genres": movie.genres,
            "cast": movie.cast,
        })

    def findByGenres(self, values):
        print("[Database] Filtering categories with values: ", values)
        return self.collections.find({"genres": {"$in": list(values)}})

    def findByCast(self, values):
        print("[Database] Filtering autores with values print: ", values)
        return self.collections.find({"cast": {"$in": list(values)}})
        
    def delete(self, movie): 
        self.collections.delete_one({"_id": movie.id})