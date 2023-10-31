<script lang="ts">
    import ContactRow from './ContactRow.svelte';
    import { contacts, events } from './pocketbase';
    import { beforeUpdate, afterUpdate, onDestroy, onMount } from 'svelte';
    import type { ContactResponse } from './pocketbase-types';
    import { createEventDispatcher } from 'svelte';
    import Event from './EventRow.svelte';
    import { userEvents } from './gin';

    let activeId = "";
    function handleMessage(event) {
        activeId = event.detail.eventId
    }

</script>

{#each $userEvents as e}
    <Event event={e} active={(activeId == e.id)} on:message={handleMessage} />
{/each}
