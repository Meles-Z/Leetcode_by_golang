<script>
  import Button from "$lib/componets/reusable/Button.svelte";
  let count = 0;

  function increment() {
    count += 1;
  }
</script>

<main class="ml-5">
  <p>{count}</p>
  <Button label="Icrement" onClick={increment} />
</main>
