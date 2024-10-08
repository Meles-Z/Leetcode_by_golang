
<style>
    section{
        width: 30rem;
        margin: auto;
    }
    label, input, textarea{
        width: 100%;
    }
</style>

<script>
    import Product from "./product.svelte";
    import Button from "./button.svelte";
  import { each } from "svelte/internal";
    let title='';
    let price=0;
    let decription='';
    let products=[];

    function setTitle(event) {
        title=event.target.value
        
    }

    function createCommerce(){
        const newProduct={
            title:title,
            price:price,
            decription:decription

        };
       products=products.concat(newProduct)

    }
</script>

<h1>{title}</h1>

<section>
    <div>
        <label for="title">Title</label>
        <input type="text" id="title" value="{title}" on:input={setTitle}>
    </div>
    <div>
        <label for="price">Price</label>
        <input type="number" name="" id="price" bind:value={price}>
    </div>
    <div>
        <label for="desc">Description</label>
       <textarea name="" id="text" rows="4" bind:value={decription}></textarea>
    </div>
    <Button on:click={createCommerce}>Buy Ecommers app</Button>
</section>
{#if products.length==0}
<p>if not product where added</p>
  {:else}
{#each products as product}

   <Product productTitle={product.title} productPrice={product.price} productDescription={product.decription}/>
    
{/each}
{/if}



