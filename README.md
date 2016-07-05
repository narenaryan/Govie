# Govie <img src="https://github.com/narenaryan/Govie/blob/master/README_files/govie.png" alt="Smiley face" height="80" width="90">

Govie is a linux command line tool for movie buffs. Govie allows you to:


* Instantly access movie pages from IMDB
* Know the rating and other details of a new movie
* Download and maintain poster collection of the movies you like.
* Flexibility to read from command line arguments.
* Ability to downloadable posters in light speed.
* Can read and process list of movies from a file

Now let me show you how tool works.

### Opening a Movie page of IMDB

![command_open](https://github.com/narenaryan/Govie/blob/master/README_files/open_movies.png)

Use <b>-o</b> flag to open a movie's IMDB page in browser.

  ```
  $ ./govie -o american-hustle gods-of-egypt
  ```
When we type the command, it takes us to the IMDB page of respective movie. Now we will see these two IMDB pages opened for us already in browser.

![command_open](https://github.com/narenaryan/Govie/blob/master/README_files/american_hustle.png)
![command_open](https://github.com/narenaryan/Govie/blob/master/README_files/gods_egypt.png)

### Getting details of a movie

Instead of launching browser, we can have a quick peek at details. Like curious to know about rating of a movie etc. We need to use <b>-d</b> option here.

  ```
  $ ./govie -d jurassic-park titanic outlander
  ```

Then output is visible on terminal itself.
![command_open](https://github.com/narenaryan/Govie/blob/master/README_files/movie_details.png)

Sometimes a movie name can create ambuiguity. So If you know the release year of the movie, we can pass it in <b>-y</b> option to get correct details.
  
  ```
  $ ./govie -d -y=2016 tarzan
  ```
  
  This gives the details about The Legend of the Tarzan (2016).

### Saving a movie poster

We can use <b>-p</b> option to save the poster of a movie in a given directory. Govie also takes multiple movies at the same time to download their posters to the given directory.

  ```
  $ ./govie -p jurassic-park titanic outlander ~/posters
  ```
  
  We should pass movies & poster save directory path to the above -p option
  Using it we can create our own poster collection in our computer. Govie tries to fetch a decent quality poster for a movie.
  
  ![command_open](https://github.com/narenaryan/Govie/blob/master/README_files/posters.png)

### Downloading and installing Govie
If you are on Ubuntu or any linux system, you can download Govie tar package from here.

[Govie tar package](https://github.com/narenaryan/Govie/blob/master/pkg/govie.tar)

Then navigate to downloaded folder and do

  ```
   $ tar -xvf govie.tar
  ```
It gives a binary executable called <b>govie</b>. Now you can use it with the above syntax of execution. But for universal accesing of executable add a ALIAS in <b>.bashrc </b> or <b>.zshrc </b>.

### Debugging

Govie creates a file in your home directory for purpose of logging and crash reporting that file is <b>govie.log</b>. Contents of that log file are used to report bugs if tool crashes.

  ```
  ~/govie.log
  ```
  ![command_open](https://github.com/narenaryan/Govie/blob/master/README_files/govie_log.png)
