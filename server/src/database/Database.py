from pymongo.mongo_client import MongoClient
from pymongo.server_api import ServerApi
from pymongo.mongo_client import MongoClient
from bson.json_util import dumps
from bson.objectid import ObjectId


class Database:
    def __init__(self):
        uri = ""
        self.client = MongoClient(uri, server_api=ServerApi("1"))
        self.database = self.client.get_database("sample_mflix")
        self.collections = self.database.get_collection("movies")

    def insert(self, movie):
        print(f"[Database] Inserting movie with title {movie.title}")

        document_to_inset = {
            "title": movie.title,
            "directors": list(movie.directors),
            "genres": list(movie.genres),
            "cast": list(movie.cast),
        }

        return self.collections.insert_one(document_to_inset)

    def update(self, movie):
        print(f"[Database] Updating movie with title {movie.title}")

        document_to_update = {"_id": ObjectId(movie.id)}
        document_data = {
            "$set": {
                "title": movie.title,
                "directors": list(movie.directors),
                "genres": list(movie.genres),
                "cast": list(movie.cast),
            }
        }

        return self.collections.update_one(document_to_update, document_data)

    def findByGenres(self, values):
        print("[Database] Filtering categories with values: ", values)
        return self.collections.find({"genres": {"$in": list(values)}})

    def findByCast(self, values):
        print("[Database] Filtering autores with values: ", values)
        return self.collections.find({"cast": {"$in": list(values)}})

    def delete(self, movie):
        print("[Database] Delete movie with id: ", movie.id)
        return self.collections.delete_one({"_id": ObjectId(movie.id)})
