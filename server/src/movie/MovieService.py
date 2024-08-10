from movie.Movies_pb2 import Movie
from movie.Movie import MovieDTO


class MovieService:
    def __init__(self, database):
        self.database = database

    def create(self, movie):
        print("[Movie Service] Executing method create()")
        movie_created = self.database.insert(movie)
        return MovieDTO.response_create(movie_created)

    def update(self, movie):
        print("[Movie Service] Executing method update()")
        return self.database.update(movie)

    def find_by_categories(self, values):
        print("[Movie Service] Executing method findByCategories()")
        movies = self.database.findByGenres(values)
        return MovieDTO.format_many_find_response(movies)

    def finc_by_actor(self, values):
        print("[Movie Service] Executing method findByAtor()")
        movies = self.database.findByCast(values)
        return MovieDTO.format_many_find_response(movies)

    def delete(self, movie):
        print("[Movie Service] Executing method delete()")
        return self.database.delete(movie)
