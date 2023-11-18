<script>
  let links = [
    {title: 'Home', path: '/'},
    {title: 'Animals', path: '/animals'},
    {title: 'Zoo', path: '/zoo'},
  ]
</script>

<nav style="margin: auto">

<ul style="display: flex; list-style: none">
  {#each links as p}
    <li style="padding: 0.5em"><a href={p.path}>{p.title}</a></li>
  {/each}
</ul>

</nav>
