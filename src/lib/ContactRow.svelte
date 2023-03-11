<script lang="ts">

    import { beforeUpdate, createEventDispatcher, onMount } from 'svelte';
    import type { ContactResponse } from './pocketbase-types';
    const dispatch = createEventDispatcher();

    export let checked: boolean;
    export let contact: ContactResponse;

    function onClick() {
        checked = !checked;
        dispatch('message', {
            checked: checked,
            id: contact.id,
        });
    }
</script>

<tr on:click={onClick}>
    <td>
        <!-- <label>
            <input bind:checked {id} type="checkbox" class="checkbox" />
        </label> -->
    </td>
    <td>
        <div class="font-bold">{contact.first_name}</div>
    </td>
    <td>
        <div class="font-bold">{contact.last_name}</div>
    </td>
    <td>
        <div class="font-bold">{contact.status}</div>
    </td>
</tr>
