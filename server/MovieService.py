from Movies_pb2 import Movie

class MovieService:
    def __init__(self, database):
        self.database = database
    
    def create(self, movie):
        print("[Movie Service] Executing method create()")
        new_movie = Movie()
        new_movie.title = movie.title
        new_movie.directors.extend(list(movie.directors))
        new_movie.genres.extend(list(movie.genres))
        new_movie.cast.extend(list(movie.cast))

        movie_created = self.database.insert(new_movie)
        movie_response = Movie()
        movie_response.id = movie_created.inserted_id
        return movie_response
    
    def update(self, movie):
        print("[Movie Service] Executing method update()")
        return self.database.update(movie)
        
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
    
    def delete(self, movie):
        print("[Movie Service] Executing method delete()")
        return self.database.delete(movie)