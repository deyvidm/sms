<script lang="ts">
  import EventForm from "$lib/EventForm.svelte";
  import { get50Contacts } from '$lib/pocketbase'
  get50Contacts()
</script>

<div class="card w-full">
  <EventForm/>
</div>
