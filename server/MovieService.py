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
            new_movie.id = str(movie["_id"])
            new_movie.plot = movie["plot"]
            new_movie.genres.extend(list(movie["genres"]))
            new_movie.runtime = movie["runtime"]
            new_movie.cast.extend(list(movie["cast"]))
            new_movie.num_mflix_comments = movie["num_mflix_comments"]
            new_movie.title = movie["title"]
            new_movie.fullplot = movie["fullplot"]
            new_movie.languages.extend(list(movie["languages"]))
            new_movie.directors.extend(list(movie["directors"]))
            new_movie.rated = movie["rated"]
            new_movie.lastupdated = movie["lastupdated"]
            new_movie.year = movie["year"]
            new_movie.countries.extend(list(movie["countries"]))
            new_movie.type = movie["type"]
            movies_list.append(new_movie)
        return movies_list
        
    def findByAtor(self, values):
        print("[Movie Service] Executing method findByAtor()")
        movies = self.database.findByCast(values)

        movies_list = []
        for movie in movies:
            new_movie = Movie()
            new_movie.id = str(movie["_id"])
            new_movie.plot = movie["plot"]
            new_movie.genres.extend(list(movie["genres"]))
            new_movie.runtime = movie["runtime"]
            new_movie.cast.extend(list(movie["cast"]))
            new_movie.num_mflix_comments = movie["num_mflix_comments"]
            new_movie.title = movie["title"]
            new_movie.fullplot = movie["fullplot"]
            new_movie.languages.extend(list(movie["languages"]))
            new_movie.directors.extend(list(movie["directors"]))
            new_movie.rated = movie["rated"]
            new_movie.lastupdated = movie["lastupdated"]
            new_movie.year = movie["year"]
            new_movie.countries.extend(list(movie["countries"]))
            new_movie.type = movie["type"]
            movies_list.append(new_movie)

        return movies_list