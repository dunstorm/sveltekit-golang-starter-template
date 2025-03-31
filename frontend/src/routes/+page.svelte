<script lang="ts">
import { Configuration, TodosApi } from '$lib/api';
import type { GithubComDunstormMySvelteGoTemplateBackendInternalModelsTodo as ApiTodo } from '$lib/api';
import { onMount } from 'svelte';

const configuration = new Configuration({
  basePath: 'http://localhost:8080'
});

const api = new TodosApi(configuration);

let todos: ApiTodo[] = [];
let newTodoTitle = '';

onMount(async () => {
  try {
    const response = await api.apiTodosGet();
    todos = response.data;
  } catch (error) {
    console.error('Failed to fetch todos:', error);
  }
});

async function addTodo() {
  if (!newTodoTitle.trim()) return;

  try {
    const response = await api.apiTodosPost({
      title: newTodoTitle,
      completed: false
    });
    todos = [...todos, response.data];
    newTodoTitle = '';
  } catch (error) {
    console.error('Failed to add todo:', error);
  }
}

async function toggleTodo(todo: ApiTodo) {
  try {
    if (todo.id !== undefined) {
      const response = await api.apiTodosIdPut(todo.id, {
        ...todo,
        completed: !todo.completed
      });
      todos = todos.map(t => t.id === todo.id ? response.data : t);
    }
  } catch (error) {
    console.error('Failed to update todo:', error);
  }
}
</script>

<main class="max-w-xl mx-auto p-8">
  <h1 class="text-center text-2xl font-bold text-gray-800 mb-6">Todo List</h1>

  <form on:submit|preventDefault={addTodo} class="flex gap-4 mb-8">
    <input
      type="text"
      bind:value={newTodoTitle}
      placeholder="Add a new todo..."
      class="flex-1 px-3 py-2 text-base border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
    />
    <button type="submit" class="px-4 py-2 text-base bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">Add</button>
  </form>

  <ul class="list-none p-0">
    {#each todos as todo (todo.id)}
      <li class="py-2 border-b border-gray-200">
        <label class="flex items-center gap-2 cursor-pointer">
          <input
            type="checkbox"
            checked={todo.completed}
            on:change={() => toggleTodo(todo)}
            class="h-5 w-5 text-blue-600 rounded"
          />
          <span class={todo.completed ? "line-through text-gray-500" : "text-gray-800"}>{todo.title}</span>
        </label>
      </li>
    {/each}
  </ul>
</main>
