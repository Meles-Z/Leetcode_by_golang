<script>
  import {goto} from '$app/navigation'
  import Test from "$lib/component/Test/test.svelte";
  import Main from "../lib/component/main.svelte";
  

  let namaste=" Hello, Svelte! {namaste}  ��"

  const person={
    test:"Namaster namaha",
    age:12,
  }
  function handleClick() {
    goto('/home')
  }
</script>

<button on:click="{handleClick}">
  Click me
</button>
<Main />
<Test {...person} />
