<div id="top"></div>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://www.linkedin.com/in/bhavyansh-jain/">
    <img src="images/me.jpg" alt="Logo" width="96" height="96">
  </a>
  <br/>
  <br/>

  <h2 align="center"> <strong> Appointy Tech Task </strong> </h2>

</div>

<hr/>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>

<hr/>

<!-- ABOUT THE PROJECT -->
## About The Project

### Task:

Develop a basic version of aInstagram.
<br/>
Required to develope an API for the system with the following specifications:

Backend controller/model/storage specifications:
- User creation/ id retrieval 
- Post creation with id and time-stamp retrieval 
  
<br/>

### Built With:

This remarkable API is built withing a day with the following language/database dependencies.

- [Golang](https://golang.org/)
- [Mongodb](https://www.mongodb.com/)

<br/>
<hr/>

<!-- GETTING STARTED -->
## Getting Started

This API can potentially link the software and application to some defined storage

### Prerequisites

- Golang installed locally
- Postman API installed locally
- MongoDB installed locally and running on port 27017

<br/>
<hr/>

<!-- USAGE -->
## Features
Features  | Details
------------ | -------------
Secure passwords  |  Stored as md5 hashes 
RFC 3339 time-stamp |  Event logging

<br/>
<hr/>

<!-- USAGE -->
## Usage
The below table specifies the url formats of the local server requests that’d be passed and the respective outputs obtained.

Endpoint | Usage
------------ | -------------
/user |  To create a new user
/users |  Displays a list of all users
/users/:id | To display a user
/post | To create a post
/posts | Displays a list of all posts
/posts/users/:id | To display all posts of a user

<br/>
<hr/>


<!-- HOW TO USE -->
## How To Use

- Make sure mongo is up and running on port 27017
- Clone this repo locally
- Open terminal  
- run this command in the project directory <br/> ``` go run main.go ```  
<br/>
- Open Postman and follow the below screenshot thread

<br/>

<div>

- Create user 1
<img src="images/create-a-user.png" alt="pic" >
- Create user 2
<img src="images/create-a-user-2.png" alt="pic" >
- Get a user
<img src="images/get-a-user.png" alt="pic" >
- Get all users
<img src="images/get-all-users.png" alt="pic" >
- Create post 1
<img src="images/create-a-post.png" alt="pic" >
- Create post 2
<img src="images/create-a-post-2.png" alt="pic" >
- Get all posts
<img src="images/get-all-posts.png" alt="pic" >
- Get all posts for a user
<img src="images/get-all-posts-for-a-user.png" alt="pic" >

</div>
