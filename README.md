# Flutter Clean Architecture Generator

This tool generates a clean architecture template for your Flutter projects, following industry best practices. It automates the creation of core folders and files for features in your project.

## Features
- Generates a clean architecture structure for Flutter projects.
- Includes Core folders (`widgets`, `error`, `theme`, `loaders`).
- Creates feature-specific folders (`data`, `domain`, `presentation`) with example files.
- Saves time and ensures consistent project structure.

---

## Usage Instructions

### Steps to Download and Run the `.exe` File

1. **Download the `.exe` File**
   - Visit the [Releases](https://github.com/yash9111/Flutter-Clean-Architecture-Template-Generator/releases/tag/go) section of this repository.
   - Download the latest `flutter_clean_arc.exe` file for your Windows system.

2. **Add the `.exe` File to System Environment Variables** *(Optional)*  
   To access the `.exe` from anywhere in the command prompt:
   - Copy the `.exe` file to a permanent location, such as `C:\Tools`.
   - Add the file location to your system's PATH:
     1. Open the **Start Menu**, search for `Environment Variables`, and select **Edit the system environment variables**.
     2. In the **System Properties** window, click **Environment Variables**.
     3. Under **System variables**, find the `Path` variable and click **Edit**.
     4. Click **New** and add the path to the folder containing the `.exe` (e.g., `C:\Tools`).
     5. Click **OK** to save and close all windows.
   - Open a new Command Prompt and type the following to verify:
     ```cmd
     flutter_clean_arc <feature_name>
     ```
   - If you see the structure created for your feature name, the setup is successful!

3. **Run the Command**
   - Open the Command Prompt (`cmd`).
   - Run  with the desired feature names as arguments:
     ```cmd
     flutter_clean_arc auth profile settings
     ```

   - This will generate the following folder structure in your Flutter project's `lib` directory:
     ```
     lib/
     ├── core/
     │   ├── widgets/
     │   ├── error/
     │   ├── theme/
     │   └── loaders/
     ├── features/
     │   ├── auth/
     │   │   ├── data/
     │   │   ├── domain/
     │   │   └── presentation/
     │   ├── profile/
     │   │   ├── data/
     │   │   ├── domain/
     │   │   └── presentation/
     │   └── settings/
     │       ├── data/
     │       ├── domain/
     │       └── presentation/
     ```

4. **Start Coding!**
   - Open the generated folder structure in your favorite IDE (e.g., VS Code or Android Studio).
   - Begin implementing your app logic using the clean architecture foundation.

---

## Requirements
- Windows Operating System.
- No additional dependencies are required to run the `.exe` file.

---

## Contributing
Feel free to fork this repository and submit pull requests to enhance the tool. Suggestions and issues can be reported [here](https://github.com/yash9111/Flutter-Clean-Architecture-Template-Generator/issues).

## License
This project is licensed under the [MIT License](LICENSE).
