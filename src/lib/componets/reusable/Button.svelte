<script lang="ts">
  export let label;
  export let onClick;
</script>

<button on:click={onClick} class="bg-blue-500 font-bold text-white py-2 px-4 hover:bg-blue-700 rounded">
  {label}
</button>
