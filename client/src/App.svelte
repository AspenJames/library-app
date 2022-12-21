<script>
  import Header from  './lib/Header.svelte';
  import Hero from    './lib/Hero.svelte';
  import Footer from  './lib/Footer.svelte';
</script>

<head>
  <title>Library App</title>
</head>
<body class="container-body">
  <div class="item-header">
    <Header homepage></Header>
  </div>
  <div class="item-hero">
    <Hero></Hero>
  </div>
  <div class="item-footer">
    <Footer></Footer>
  </div>
</body>

<style>
  .container-body{
    --body-width: 1280px;
    display: grid;
    grid-template-columns: var(--body-width);
    grid-template-rows: [first] 50px [line2] 600px [line3] 200px [last];

    background-color: white;
    font-family: "Montserrat", Sans-Serif;
    width: var(--body-width);
    height: auto;
  }
  .item-header{
    grid-column-start: 1;
    grid-row-start: 1;
  }
  .item-hero{
    grid-column-start:1;
    grid-column-end:2;
    grid-row-start:2;
    grid-row-end:3;
    width: 100%;
    height: 100%;
  }
  .item-footer{
    grid-column-start: 1;
    grid-column-end: 2 ;
    grid-row-start: 3;
    grid-row-end: 4;
    object-fit: cover;
  }
</style>