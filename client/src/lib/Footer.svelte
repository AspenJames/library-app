<script>
    import Logo from './Logo.svelte'
</script>

<div class="container">
    <div class="subcontainer1">
        <div class="hyperlink1">
            <p>Mobile App</p>
        </div>
        <div class="hyperlink2">
            <p>About</p>
        </div>
        <div class="logo">
            <Logo></Logo>
        </div>
        <div class="hyperlink3">
            <p>Blog</p>
        </div>
        <div class="hyperlink4">
            <p>Resources</p>
        </div>
    </div>
    <div class="subcontainer2"></div>
    <div class="subcontainer3"> <p>Fake copyright Library inc. 2022</p></div>
</div>

<style>
    .container{
        --footer-height: 200px;
        --footer-width: 1280px;

        height: var(--footer-height);
        width: var(--footer-width);
        background-color: white;
        color: black;
        display: grid;
        grid-template-columns: auto;
        grid-template-rows: 80px 1px auto;
    }
    .subcontainer1{
        grid-column-start: 1;
        grid-row-start: 1;
        display: grid;
        grid-template-columns: repeat(5, 20%);
        grid-template-rows: auto;
        margin-top: 10px;
        margin-left: 80px;
        margin-right: 80px;
        min-height: 0;
    }
    .subcontainer2{
        background-color: rgb(47, 47, 47);
        margin-left: 150px;
        margin-right: 150px;
    }
    .subcontainer3{
        font-size: 11px;
        color: grey;
        margin-top: 40px;
    }
    .hyperlink1{
        grid-column-start: 1;
        grid-column-end: 2;
        grid-row-start: 1;
        object-fit: cover;
        justify-self: center;
        font-size: 12px;
        font-weight: bold;
    }
    .hyperlink2{
        grid-column-start: 2;
        grid-column-end: 3;
        grid-row-start: 1;
        object-fit: cover;
        justify-self: center;
        font-size: 12px;
        font-weight: bold;
    }
    .logo{
        grid-column-start: 3;
        grid-column-end: 4;
        grid-row-start: 1;
        justify-self: center;
        margin-left: 20px;
        margin-right: 20px;
        max-height: 100%;
        max-width: 100%;
    }
    .hyperlink3{
        grid-column-start: 4;
        grid-column-end: 5;
        grid-row-start: 1;
        object-fit: cover;
        justify-self: center;
        font-size: 12px;
        font-weight: bold;
    }
    .hyperlink4{
        grid-column-start: 5;
        grid-column-end: 6;
        grid-row-start: 1;
        object-fit: cover;
        justify-self: center;
        font-size: 12px;
        font-weight: bold;
    }
</style>