<script lang="ts">

    import { beforeUpdate, createEventDispatcher, onMount } from 'svelte';
    import type { ContactResponse } from './pocketbase-types';
    import type { Contact } from './gripes';
    const dispatch = createEventDispatcher();

    export let checked: boolean;
    export let contact: Contact;

    function onClick() {
        checked = !checked;
        dispatch('message', {
            checked: checked,
            id: contact.id,
        });
    }
</script>

<tr on:click={onClick}>
    <td hidden={true} >
        <label>
            <input bind:checked id={contact.id} type="checkbox" class="checkbox"  />
        </label>
    </td>
    <td>
        <div class="font-bold">{contact.first_name}</div>
    </td>
    <td>
        <div class="font-bold">{contact.last_name}</div>
    </td>
    <td>
        <div class="font-bold">{contact.phone}</div>
    </td>
</tr>
