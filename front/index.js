console.log("hello golang and vue!!");

fetch("/messages")
  .then(res => res.json())
  .then(data => console.log(data));
