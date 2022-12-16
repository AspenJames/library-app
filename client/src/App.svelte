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
    grid-template-rows: [first] 50px [line2] auto [line3] 200px [last];

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




<!--
<Logo></Logo>
-->

<!--
<Lib_filter_menu active>All</Lib_filter_menu>
<Lib_filter_menu>Read</Lib_filter_menu>
<Lib_filter_menu>Unread</Lib_filter_menu>
-->


<!--
<Button>Log In</Button>

<Button>Log Out</Button>

<Button orange_button>Sign Up</Button>

<Button wide_orange_button>Create Account</Button>

<Button add_book_button>+</Button>

<Row_tag>Unread</Row_tag>

<Row_tag read_tag>Read</Row_tag>

<Row></Row>

<Logo></Logo> 

<Logo large_logo></Logo>


<Card></Card>

<Card homepage_card></Card>
-->


