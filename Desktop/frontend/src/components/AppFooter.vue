<template>
  <v-footer height="40" app color="transparent">
    <a
      v-for="item in items"
      :key="item.title"
      :href="item.href"
      :title="item.title"
      class="d-inline-block mx-2 social-link"
      rel="noopener noreferrer"
      target="_blank"
    >
      <v-icon
        :icon="item.icon"
        :size="item.icon === '$vuetify' ? 24 : 16"
      />
    </a>

    <div
      class="text-caption text-disabled"
      style="position: absolute; right: 16px;"
    >
      &copy; 2025 <span class="d-none d-sm-inline-block">Algebra</span>
    </div>
  </v-footer>
</template>

<script setup lang="ts">
  const items = [
    {
      title: 'Me',
      icon: `mdi-linkedin`,
      href: 'https://www.linkedin.com/in/lorena-kova%C4%8Di%C4%87/',
    },
    {
      title: 'Project github',
      icon: `mdi-github`,
      href: 'https://github.com/Tiitmouse/iis_project',
    },
  ]
</script>