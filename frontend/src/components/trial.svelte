<script lang="ts">
	import { Label } from 'flowbite-svelte';
	let file: FileList;
	let data: any[] = []; 
	let loading = false; // Declare a new reactive variable for the loading state

	const uploadFile = async () => {
		loading = true; // Set loading to true before fetching data
		const formData = new FormData();
		formData.append('file', file[0]);

		try {
			const res = await fetch('http://127.0.0.1:3000/upload', {
				method: 'POST',
				body: formData
			});
			data = await res.json(); 
			console.log('File uploaded successfully', res);
		} catch (err) {
			console.error('Error uploading file', err);
		} finally {
			loading = false; // Set loading to false when done
		}
	};

	const convertToCSV = (objArray) => {
    const array = typeof objArray != 'object' ? JSON.parse(objArray) : objArray;
    let str = '';

    for (let i = 0; i < array.length; i++) {
        let line = '';
        for (let index in array[i]) {
            if (line !== '') line += ','
            line += `"${array[i][index]}"`;
        }
        str += line + '\r\n';
    }
    return str;
}

    const exportCSVFile = () => {
        let csv = convertToCSV(data);
        let exportedFilename = 'export.csv';

        let blob = new Blob([csv], { type: 'text/csv;charset=utf-8;' });
        if (navigator.msSaveBlob) { // IE 10+
            navigator.msSaveBlob(blob, exportedFilename);
        } else {
            let link = document.createElement("a");
            if (link.download !== undefined) { // feature detection
                // Browsers that support HTML5 download attribute
                let url = URL.createObjectURL(blob);
                link.setAttribute("href", url);
                link.setAttribute("download", exportedFilename);
                link.style.visibility = 'hidden';
                document.body.appendChild(link);
                link.click();
                document.body.removeChild(link);
            }
        }
    }
</script>

<div class="p-8">
	<Label class="space-y-2 mb-2">
		<span>Upload file</span>
		<form on:submit|preventDefault={uploadFile}>
			<input type="file" accept=".csv, .xlsx" bind:files={file} />
			<button type="submit">Upload</button>
		</form>
	</Label>
	{#if loading}
		<p>Loading...</p>
		<!-- Create a table with headers and rows based on the server response -->
		<!-- ... rest of your table code ... -->
	{/if}
	{#if data.length > 0}
		<!-- Create a table with headers and rows based on the server response -->
		<div class="px-4 sm:px-6 lg:px-8">
			<div class="sm:flex sm:items-center">
			  <div class="sm:flex-auto">
				<h1 class="text-base font-semibold leading-6 text-gray-900">Data from your list</h1>
				<p class="mt-2 text-sm text-gray-700">Download your list to the right</p>
			  </div>
			  <div class="mt-4 sm:ml-16 sm:mt-0 sm:flex-none">
				<button type="button" on:click={exportCSVFile} class="block rounded-md bg-indigo-600 px-3 py-2 text-center text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Export</button>
			</div>
			</div>
			<div class="mt-8 flow-root">
			  <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
				<div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
				  <table class="min-w-full divide-y divide-gray-300">
					<thead>
						<tr>
						
						{#each Object.keys(data[0]) as key}
						<th class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">{key}</th>
					{/each}
						<th scope="col" class="relative py-3.5 pl-3 pr-4 sm:pr-0">
						  <span class="sr-only">Edit</span>
						</th>
					  </tr>
					</thead>
					<tbody class="divide-y divide-gray-200">
						
						{#each data as row}
					<tr>
						{#each Object.values(row) as value}
							<td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{value}</td>
						{/each}
					</tr>
				{/each}
		  
					  <!-- More people... -->
					</tbody>
				  </table>
				</div>
			  </div>
			</div>
		  </div>
		  
	
	{/if}
</div>
