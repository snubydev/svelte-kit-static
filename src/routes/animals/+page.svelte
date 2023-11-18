<script>
  export let data = { products: [], next: 20 }

  function on_loaded ({detail}) {
    data.products = [...data.products, ...detail.products]
    data.next = detail.next
  }

</script>

<h2>Animals</h2>

{#each data.products as a}
  <div>{a}</div>
{/each}