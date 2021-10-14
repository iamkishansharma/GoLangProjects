/*
Part of exercise file for go lang course at
https://web.learncodeonline.in
*/

const express = require("express");
const app = express();
const port = 8000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).send(webPage);
});

app.get("/get", (req, res) => {
  res.status(200).json({
    message: "Hey, you are here because of you're learning Go from LCO.",
  });
});

app.get("/get/users", (req, res) => {
  res.status(200).json(usersData);
});

app.post("/post", (req, res) => {
  let myJson = req.body; // your JSON

  res.status(200).send(myJson);
  //   usersData.push(myJson);
});

app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});

const webPage = `
<!DOCTYPE html>
<html>
    <head>
        <title>GoLang Web Server | Kishan</title>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
        <link href="https://fonts.googleapis.com/css2?family=Oleo+Script&display=swap" rel="stylesheet" />

        <style>
            .header {
                font-family: "Oleo Script", cursive;
                padding: 10px;
                color: white;
                background-color: #5a20cb;
                text-align: center;
                font-size: 50px;
                border-radius: 5px;
                text-shadow: 2px 2px 4px #000000;
                box-shadow: 1px 2px 10px #000;
            }
            .button {
                padding: 15px;
                background-color: #242b2e;
                text-align: center;
                font-size: 20px;
                font-weight: bold;
                box-shadow: 2px 2px 10px #000;
                border-radius: 5px;
            }
            a {
                text-decoration: none;
                width: 160px;
                color: white;display:block;
            }
            .button:hover {
                background-color: purple;
                transition: 0.5s;
                transform: scale(1);
                background-color: #5a20cb;
            }
        </style>
    </head>
    <body style="background-color: purple; color: white;">
        <div class="header">Kishan Kumar Sharma</div>
        <center>
            <h1 style="align: center;">Welcome to GoLang Web Server.</h1>
            <h2>GET | POST | PUT | DELETE</h2>
            <h3>built for Go lang course on LCO</h3>
            <a href="https://github.com/iamkishansharma" target="_blank"><div class="button">GitHub</div></a>
        </center>
    </body>
</html>

`;

const usersData = [
  {
    id: 1,
    name: "John Wick",
    email: "johnwick3@film.com",
    dob: "1964-09-02",
    image:
      "https://sm.ign.com/ign_ap/cover/j/john-wick-/john-wick-chapter-4_z6zv.jpg",
    age: 57,
  },
  {
    id: 2,
    name: "Peter Parker",
    email: "peter@spiderman.com",
    dob: "1975-06-27",
    image:
      "https://www.chinadaily.com.cn/entertainment/att/site1/20070306/xin_35030406155324208281.jpg",
    age: 46,
  },
  {
    id: 3,
    name: "Maha Nayak",
    email: "rajesh@hamal.hit",
    dob: "1964-06-09",
    image: "https://rajeshhamal.com.np/wp-content/uploads/2018/08/profile.jpg",
    age: 57,
  },
  {
    id: 4,
    name: "Natasha Romanoff",
    email: "natasha.romanoff@cmail.com",
    dob: "1984-11-22",
    image: "https://i.quotev.com/img/q/u/18/4/1/kysqanlhxe.jpg",
    age: 36,
  },
  {
    id: 5,
    name: "The Rock",
    email: "therock@nomail.com",
    dob: "1972-05-02",
    image:
      "https://upload.wikimedia.org/wikipedia/commons/f/f1/Dwayne_Johnson_2%2C_2013.jpg",
    age: 49,
  },
  {
    id: 6,
    name: "Kishan Sharma",
    email: "skishan781@gmail.com",
    dob: "1998-08-25",
    image: "https://avatars.githubusercontent.com/u/36340195",
    age: 23,
  },
  {
    id: 7,
    name: "Harley Quinn",
    email: "harley.quinn@official.com",
    dob: "1990-07-02",
    image:
      "https://i1.wp.com/www.cinedergi.com/wp-content/uploads/2020/07/5c8d8a7645d2a09e0099d447.jpg?resize=678%2C482",
    age: 31,
  },
  {
    id: 8,
    name: "Sundar Pichai",
    email: "sundarpichai@google.com",
    dob: "1972-06-10",
    image:
      "https://pbs.twimg.com/profile_images/864282616597405701/M-FEJMZ0_400x400.jpg",
    age: 49,
  },
  {
    id: 9,
    name: "Capt.Jack Sparrow",
    email: "jacksparrow@email.com",
    dob: "1963-06-09",
    image:
      "https://images.saatchiart.com/saatchi/1363531/art/6693441/5763105-NCQCYGJL-7.jpg",
    age: 58,
  },
  {
    id: 10,
    name: "Tony Stark",
    email: "tony.stark@ironman.com",
    dob: "1965-04-04",
    image:
      "https://i.pinimg.com/originals/8e/21/29/8e2129f44804db65316ed3db92cf8552.jpg",
    age: 56,
  },
  {
    id: 11,
    name: "Brian O'Conner",
    email: "brian@fandf.com",
    dob: "1973-09-12",
    image:
      "https://www.soapoperadigest.com/wp-content/uploads/files/inline/bergman38591.jpg",
    age: 48,
  },
];
