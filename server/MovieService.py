from Database import Database
from Movies_pb2 import Response, Movie
from bson.json_util import dumps
import logging

class MovieService:
    def __init__(self, database):
        self.database = database
    
    def create(self, movie):
        print("[Movie Service] Executing method create()")
        movie_created = self.database.insert(movie)

    def findByCategories(self, values):
        print("[Movie Service] Executing method findByCategories()")
        movies = self.database.findByGenres(values)

        movies_list = []
        for movie in movies:
            new_movie = Movie()
            new_movie.title = movie["title"]
            movies_list.append(new_movie)

        return movies_list
        
    def findByAtor(self, values):
        print("[Movie Service] Executing method findByAtor()")
        movies = self.database.findByCast(values)

        movies_list = []
        for movie in movies:
            new_movie = Movie()
            print(movie["_id"])
            print(movie["plot"])
            print(movie["type"])
            print(movie)
            new_movie.title = movie["title"]
            new_movie.plot = movie["plot"]
            movies_list.append(new_movie)

        return movies_list