<script lang="ts">
  import Counter from './lib/Counter.svelte'
  
  const tableData = $state([
  { position: '0', name: 'Iron', symbol: 'Fe', atomic_no: '26' },
  { position: '1', name: 'Rhodium', symbol: 'Rh', atomic_no: '45' },
  { position: '2', name: 'Iodine', symbol: 'I', atomic_no: '53' },
  { position: '3', name: 'Radon', symbol: 'Rn', atomic_no: '86' },
  { position: '4', name: 'Technetium', symbol: 'Tc', atomic_no: '43' }
]);

const handleClick = ()  => {
	tableData.push(
		{ position: '4', name: 'Technetium', symbol: 'Tc', atomic_no: '43' }
	)
}
</script>

<main>
<button onclick={handleClick} type="button" class="btn preset-filled">add +</button>



<div class="table-wrap">
	<table class="table caption-bottom">
	  <caption class="pt-4">A list of elements from the periodic table.</caption>
	  <thead>
		<tr>
		  <th>Position</th>
		  <th>Symbol</th>
		  <th>Name</th>
		  <th class="!text-right">Weight</th>
		</tr>
	  </thead>
	  <tbody class="[&>tr]:hover:preset-tonal-primary">
		{#each tableData as row}

		<tr>
			<td>{row.position}</td>
			<td>{row.symbol}</td>
			<td>{row.name}</td>
			<td class="text-right">{row.atomic_no}</td>
		  </tr>
		{/each}

	  </tbody>
	  <tfoot>
		<tr>
		  <td colspan="3">Total</td>
		  <td class="text-right">{tableData.length} Elements</td>
		</tr>
	  </tfoot>
	</table>
  </div>

</main>

<style>
</style>
