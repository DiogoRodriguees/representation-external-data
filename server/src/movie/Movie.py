from movie.Movies_pb2 import Movie


class MovieDTO:
    @staticmethod
    def response_create(movie_created):
        movie_response = Movie()
        movie_response.id = movie_created.inserted_id
        return movie_response

    @staticmethod
    def format_find_response(movie):
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

        return new_movie

    @staticmethod
    def format_many_find_response(movies):
        movies_list = []
        for movie in movies:
            new_movie = MovieDTO.format_find_response(movie)
            movies_list.append(new_movie)
        return movies_list
