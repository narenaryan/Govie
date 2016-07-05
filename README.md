# Govie <img src="https://github.com/narenaryan/Govie/blob/master/README_files/govie.png" alt="Smiley face" height="80" width="90">

Govie is a linux command line tool for movie buffs. Govie allows you to:


* Instantly access movie pages from IMDB
* Know the rating and other details of a new movie
* Download and maintain poster collection of the movies you like.
* Flexibility to read from command line arguments.
* Ability to create poster collection from a file.

Now let me show you how tool works.

### Opening a Movie page of IMDB

![command_open](https://github.com/narenaryan/Govie/blob/master/README_files/open_movies.png)

Use <b>-o</b> flag to open a movie's IMDB page in browser.

  $ ./govie -o american-hustle gods-of-egypt

When we type the command, it takes us to the IMDB page of respective movie. Now we will see these two IMDB pages opened for us already in browser.

![command_open](https://github.com/narenaryan/Govie/blob/master/README_files/american_hustle.png)
![command_open](https://github.com/narenaryan/Govie/blob/master/README_files/gods_egypt.png)

### Getting details of a movie

Instead of launching browser, we can have a quick peek at details. Like curious to know about rating of a movie etc. We need to use <b>-d</b> flag here.

  ./govie -d jurassic-park titanic outlander

Then output is visible on terminal itself.
![command_open](https://github.com/narenaryan/Govie/blob/master/README_files/movie_details.png)


