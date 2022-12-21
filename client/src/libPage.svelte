<script>
    import Header from  './lib/Header.svelte';
    import Lib_table from './lib/Library_table.svelte'
    import Footer from  './lib/Footer.svelte';

  
  </script>
  
  <head>
    <title>Library App - Personal Library</title>
  </head>
  <body class="container-body">
    <div class="item-header">
      <Header></Header>
    </div>
    <div class="item-lib-table">
      <Lib_table></Lib_table>
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
    .item-lib-table{
      position: relative;
      grid-column-start:1;
      grid-column-end:2;
      grid-row-start:2;
      grid-row-end:3;
      width: 100%;
      height: 100%;;
    }
    .item-footer{
      grid-column-start: 1;
      grid-column-end: 2 ;
      grid-row-start: 3;
      grid-row-end: 4;
      object-fit: cover;
      margin-top: 200px;
    }
  </style>