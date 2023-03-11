<script lang="ts">
    import ContactRow from './ContactRow.svelte';
    import { contacts, events } from './pocketbase';
    import { beforeUpdate, afterUpdate, onDestroy, onMount } from 'svelte';
    import type { ContactResponse } from './pocketbase-types';
    import { createEventDispatcher } from 'svelte';
    import EventRow from './EventRow.svelte';

    let activeId = "";
    function handleMessage(event) {
        activeId = event.detail.eventId
        console.log(activeId)
    }

</script>

{#each $events as e}
    <EventRow event={e} active={(activeId == e.id)} on:message={handleMessage} />
{/each}
