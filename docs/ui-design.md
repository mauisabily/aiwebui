# UI Design Specification

## Overview
This document describes the UI design for the AI WebUI application, built with Vue.js for a modern reactive interface similar to OpenWebUI.

## Design Principles
- Clean, modern interface with dark/light theme support
- Responsive design for desktop and tablet use
- Intuitive navigation and workflow
- Real-time feedback during operations
- Accessibility considerations

## Color Scheme
- Primary: #6366F1 (Indigo)
- Secondary: #8B5CF6 (Violet)
- Background: #111827 (Dark) / #F9FAFB (Light)
- Text: #F3F4F6 (Dark bg) / #1F2937 (Light bg)
- Success: #10B981 (Emerald)
- Warning: #F59E0B (Amber)
- Error: #EF4444 (Red)

## Layout Structure

```
+-------------------------------------------------------------+
| Header                                                      |
| [Logo] [App Name]        [Model Selector] [Settings] [User] |
+-------------------------------------------------------------+
| Sidebar |                                                   |
|         |                                                   |
| Chats   |                                                   |
| +------+|                                                   |
| |Chat 1||                                                   |
| |Chat 2||                      Main Content                 |
| |Chat 3||                                                   |
| +------+|                                                   |
|         |                                                   |
| Knowledge Bases                                             |
| +-----------------+                                         |
| |Knowledge Base 1 |                                         |
| |Knowledge Base 2 |                                         |
| +-----------------+                                         |
+-------------------------------------------------------------+
| Footer                                                      |
| [Status] [Progress]                                         |
+-------------------------------------------------------------+
```

## Component Breakdown

### 1. Header Component
- Application logo and name
- Model selector dropdown
- Settings button/modal
- User profile menu

### 2. Sidebar Component
#### Chat Section
- List of conversation threads
- Ability to create new chat
- Search/filter conversations
- Delete/archive conversations

#### Knowledge Base Section
- List of knowledge bases
- Create/manage knowledge bases
- Import/upload documents
- View knowledge base statistics

### 3. Main Content Area
#### Chat Interface
- Message bubbles for user and AI
- Input area with text field and send button
- Streaming response visualization
- Copy message functionality
- Regenerate response option
- Like/dislike feedback

#### Knowledge Base Management
- Document list with titles and metadata
- Upload/import interface
- Search/indexing status
- Delete/edit documents

### 4. Modals and Overlays
#### Settings Modal
- Theme selection (dark/light)
- Default model selection
- Auto-save preferences
- Language settings

#### Model Management Modal
- List of available models
- Download/remove models
- Model information/details

#### Knowledge Base Creation Modal
- Name and description fields
- Privacy settings
- Initial document import

## Vue Component Hierarchy

```
App.vue
├── Header.vue
│   ├── Logo.vue
│   ├── ModelSelector.vue
│   ├── SettingsButton.vue
│   └── UserProfile.vue
├── Sidebar.vue
│   ├── ChatList.vue
│   │   ├── ChatItem.vue
│   │   └── NewChatButton.vue
│   └── KnowledgeBaseList.vue
│       ├── KnowledgeBaseItem.vue
│       └── NewKnowledgeBaseButton.vue
├── MainContent.vue
│   ├── ChatView.vue
│   │   ├── MessageList.vue
│   │   │   ├── UserMessage.vue
│   │   │   └── AIMessage.vue
│   │   ├── MessageInput.vue
│   │   └── ChatActions.vue
│   └── KnowledgeBaseView.vue
│       ├── DocumentList.vue
│       │   ├── DocumentItem.vue
│       │   └── UploadDocument.vue
│       └── KnowledgeBaseActions.vue
├── Modals.vue
│   ├── SettingsModal.vue
│   ├── ModelManagerModal.vue
│   └── KnowledgeBaseModal.vue
└── Footer.vue
```

## UI State Management

### Vuex Store Modules
1. **Chat Module**
   - Current conversation
   - Message history
   - Loading states
   - Streaming status

2. **Model Module**
   - Available models
   - Selected model
   - Model capabilities

3. **KnowledgeBase Module**
   - Knowledge bases list
   - Selected knowledge base
   - Documents
   - Search results

4. **Settings Module**
   - Theme preference
   - User preferences
   - Application settings

## Responsive Design

### Desktop (> 1024px)
- Full sidebar visible
- Three-column layout (sidebar, main content, optional right panel)

### Tablet (768px - 1024px)
- Collapsible sidebar
- Two-column layout
- Stacked modals

### Mobile (< 768px)
- Bottom navigation
- Full-screen modals
- Simplified interface

## Accessibility Features

- Keyboard navigation support
- Screen reader compatibility
- Proper contrast ratios
- Focus indicators
- ARIA labels for interactive elements

## Performance Considerations

- Virtual scrolling for long message lists
- Lazy loading of images/documents
- Efficient rendering with Vue's reactivity system
- Debounced search inputs
- Caching of frequently accessed data

## CSS Framework

We'll use Tailwind CSS for styling with the following customizations:

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  :root {
    --color-primary: #6366F1;
    --color-secondary: #8B5CF6;
    --color-success: #10B981;
    --color-warning: #F59E0B;
    --color-error: #EF4444;
  }
  
  .dark {
    --bg-primary: #111827;
    --bg-secondary: #1F2937;
    --text-primary: #F3F4F6;
    --text-secondary: #D1D5DB;
  }
  
  .light {
    --bg-primary: #F9FAFB;
    --bg-secondary: #F3F4F6;
    --text-primary: #1F2937;
    --text-secondary: #6B7280;
  }
}

@layer components {
  .message-user {
    @apply bg-indigo-500 text-white rounded-l-xl rounded-tr-xl py-3 px-4 max-w-[80%] self-end;
  }
  
  .message-ai {
    @apply bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-gray-200 rounded-r-xl rounded-tl-xl py-3 px-4 max-w-[80%] self-start;
  }
  
  .btn-primary {
    @apply bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2 px-4 rounded-lg transition duration-200;
  }
  
  .btn-secondary {
    @apply bg-gray-200 dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-600 text-gray-800 dark:text-gray-200 font-medium py-2 px-4 rounded-lg transition duration-200;
  }
}
```

## Icon System

We'll use Heroicons (Vue components) for consistent iconography:

- Chat icons: ChatBubbleLeftRightIcon, PlusIcon, TrashIcon
- Knowledge base icons: BookOpenIcon, DocumentPlusIcon, MagnifyingGlassIcon
- Settings icons: Cog6ToothIcon, MoonIcon, SunIcon
- Model icons: CpuChipIcon, ArrowDownTrayIcon
- Action icons: PaperAirplaneIcon, StopIcon, ArrowsClockwiseIcon

## Animation and Transitions

- Smooth transitions between views
- Typing indicators for AI responses
- Fade-in animations for new messages
- Progress bars for long operations
- Micro-interactions for button clicks

## Internationalization Support

Structure for supporting multiple languages:
- English (default)
- Malay (Bahasa Melayu)
- Chinese (Simplified)

## Error Handling UI

- Inline error messages
- Toast notifications for async operations
- Empty state illustrations
- Retry mechanisms for failed operations

This UI design provides a comprehensive foundation for building a modern, feature-rich AI WebUI application that maintains the familiarity of OpenWebUI while being streamlined for local model usage.