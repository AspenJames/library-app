<script>
    import Logo from './Logo.svelte'
    import Link from './Link.svelte'
    import Cprt from './Copyright.svelte'
</script>

<div class="container">
    <div class="links-and-logo">
        <div class="hyperlink1">
            <Link link="#home">Mobile App</Link>
        </div>
        <div class="hyperlink2">
            <Link link="#home">About</Link>
        </div>
        <div class="logo">
            <Logo></Logo>
        </div>
        <div class="hyperlink3">
            <Link link="#home">Blog</Link>
        </div>
        <div class="hyperlink4">
            <Link link="#home">Resources</Link>
        </div>
    </div>
    <div class="subcontainer2"></div>
    <div class="subcontainer3"><Cprt></Cprt>></div>
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
        grid-template-rows: 60px 1px auto;
    }
    .links-and-logo{
        grid-column-start: 1;
        grid-row-start: 1;
        grid-row-end: 2;
        display: flex;
        flex-direction: row;
        justify-content: space-around;
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
        align-self: center;
        font-size: 12px;
        font-weight: bold;
    }
    .hyperlink2{
        grid-column-start: 2;
        grid-column-end: 3;
        grid-row-start: 1;
        object-fit: cover;
        justify-self: center;
        align-self: center;
        font-size: 12px;
        font-weight: bold;
    }
    .logo{
        grid-column-start: 3;
        grid-column-end: 4;
        grid-row-start: 1;
        object-fit: cover;
        justify-self: center;
        align-self: center;
        margin-left: 0px;
        margin-right: 0px;
    }
    .hyperlink3{
        grid-column-start: 4;
        grid-column-end: 5;
        grid-row-start: 1;
        object-fit: cover;
        justify-self: center;
        align-self: center;
        font-size: 12px;
        font-weight: bold;
    }
    .hyperlink4{
        grid-column-start: 5;
        grid-column-end: 6;
        grid-row-start: 1;
        object-fit: cover;
        justify-self: center;
        align-self: center;
        font-size: 12px;
        font-weight: bold;
    }
</style>