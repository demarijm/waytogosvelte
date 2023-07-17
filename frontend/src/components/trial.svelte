<script lang="ts">
	import { Label } from 'flowbite-svelte'
	let file: FileList

	const uploadFile = async () => {
		const formData = new FormData();
		formData.append('file', file[0]);

		try {
			const res = await fetch('http://127.0.0.1:3000/upload', {
				method: 'POST',
				body: formData
			});
			console.log('File uploaded successfully', res);
		} catch (err) {
			console.error('Error uploading file', err);
		}
	};
</script>

<div class="p-8">
	<Label class="space-y-2 mb-2">
		<span>Upload file</span>
		<form on:submit|preventDefault={uploadFile}>
			<input type="file" accept=".csv, .xlsx" bind:files={file} />
			<button type="submit">Upload</button>
		</form>
	</Label>
</div>
