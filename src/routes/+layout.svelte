<script>
  import Nav from '../components/Nav.svelte'

</script>

<div style="padding: 1rem; background-color: darkolivegreen; color: white">
  <h2>Static App</h2>
</div>

<div style="background-color: #d1e1b3; display: flex">
  <Nav />
</div>

<main style="padding: 1rem">
  <slot />
</main>

<style>
  :global(*) {
    margin: 0;

  }
</style>