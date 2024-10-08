
<script>
    import Button from "./button.svelte";
    export let productTitle;
    export let productPrice;
    export let productDescription  
</script>

<style>
    div{
        margin: 1rem;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26) ;
        padding: 1 rem;
        border-radius: 5px;
        background: white;
        
        }
        h1{
            font-size: 1.25rem;
            margin: 0.25rem 0;
        }
        h2{
            font-size: 1 rem;
            margin: 0.25rem 0;
            color: black;

        }
        p{
            margin: 0.25rem 0;
        }
       
</style>


<div>
    <h1>{productTitle}</h1>
    <h2>{productPrice} </h2>
    <p>{productDescription}</p>
    <Button>Add Cart</Button>
    
</div>