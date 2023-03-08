<script lang="ts">
  import EventForm from "$lib/EventForm.svelte";
    import { get50Contacts } from "$lib/pocketbase";
    import { onMount } from "svelte";
</script>

<div class="card w-full">
  <EventForm/>
</div>
