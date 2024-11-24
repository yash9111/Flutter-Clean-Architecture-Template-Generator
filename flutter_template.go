package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func createFolderIfNotExist(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create folder: %s", path)
		}
	}
	return nil
}

func createFileWithContent(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %s", filePath)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %s", filePath)
	}

	return nil
}

func createCoreStructure() error {
	coreFolders := []string{
		"core/widgets",
		"core/error",
		"core/theme",
		"core/loaders",
	}

	for _, folder := range coreFolders {
		err := createFolderIfNotExist(filepath.Join("lib", folder))
		if err != nil {
			return err
		}
	}

	errorFile := "lib/core/error/server_exception.dart"
	errorContent := `class ServerException implements Exception {
  final String message;
  
  ServerException({required this.message});
}
`
	err := createFileWithContent(errorFile, errorContent)
	if err != nil {
		return err
	}

	loaderFile := "lib/core/loaders/custom_loader.dart"
	loaderContent := `import 'package:flutter/material.dart';

class CustomLoader extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: CircularProgressIndicator(),
    );
  }
}
`
	err = createFileWithContent(loaderFile, loaderContent)
	if err != nil {
		return err
	}

	themeFile := "lib/core/theme/theme.dart"
	themeContent := `import 'package:flutter/material.dart';

class AppTheme {
  static ThemeData lightTheme = ThemeData(
    primarySwatch: Colors.blue,
    brightness: Brightness.light,
  );

  static ThemeData darkTheme = ThemeData(
    primarySwatch: Colors.blue,
    brightness: Brightness.dark,
  );
}
`
	err = createFileWithContent(themeFile, themeContent)
	if err != nil {
		return err
	}

	return nil
}

func createFeatureStructure(features []string) error {
	for _, feature := range features {
		featureFolder := filepath.Join("lib", "features", feature)

		err := createFolderIfNotExist(featureFolder)
		if err != nil {
			return err
		}

		dataFolder := filepath.Join(featureFolder, "data")
		domainFolder := filepath.Join(featureFolder, "domain")
		presentationFolder := filepath.Join(featureFolder, "presentation")

		err = createFolderIfNotExist(dataFolder)
		if err != nil {
			return err
		}

		err = createFolderIfNotExist(domainFolder)
		if err != nil {
			return err
		}

		err = createFolderIfNotExist(presentationFolder)
		if err != nil {
			return err
		}

		dataModelsFolder := filepath.Join(dataFolder, "models")
		dataReposFolder := filepath.Join(dataFolder, "repos")
		dataSourcesFolder := filepath.Join(dataFolder, "sources")

		err = createFolderIfNotExist(dataModelsFolder)
		if err != nil {
			return err
		}

		err = createFolderIfNotExist(dataReposFolder)
		if err != nil {
			return err
		}

		err = createFolderIfNotExist(dataSourcesFolder)
		if err != nil {
			return err
		}

		modelFile := filepath.Join(dataModelsFolder, feature+"_model.dart")
		modelContent := fmt.Sprintf(`
		import '../../domain/entities/%s_entity.dart';
		class %sModel extends %sEntity{}
`, strings.Title(feature), strings.Title(feature), strings.Title(feature))

		err = createFileWithContent(modelFile, modelContent)
		if err != nil {
			return err
		}

		repoFile := filepath.Join(dataReposFolder, feature+"_repo_impl.dart")
		repoContent := fmt.Sprintf(`
		import '../../domain/repo/%s_repo.dart';
		class %sRepoImpl implements %sRepo{}
`, strings.Title(feature), strings.Title(feature), strings.Title(feature))

		err = createFileWithContent(repoFile, repoContent)
		if err != nil {
			return err
		}

		remoteSourceFile := filepath.Join(dataSourcesFolder, feature+"_remote_datasource.dart")
		remoteSourceContent := fmt.Sprintf(`
		abstract class %sRemoteDataSource {}
		class %sRemoteDataSourceImpl implements %sRemoteDataSource{}
`, strings.Title(feature), strings.Title(feature), strings.Title(feature))

		err = createFileWithContent(remoteSourceFile, remoteSourceContent)
		if err != nil {
			return err
		}

		localSourceFile := filepath.Join(dataSourcesFolder, feature+"_local_datasource.dart")
		localSourceContent := fmt.Sprintf(`
		abstract class %sLocalDataSource {}
		class %sLocalDataSourceImpl implements %sLocalDataSource{}
`, strings.Title(feature), strings.Title(feature), strings.Title(feature))

		err = createFileWithContent(localSourceFile, localSourceContent)
		if err != nil {
			return err
		}

		domainRepoFolder := filepath.Join(domainFolder, "repo")
		domainEntitiesFolder := filepath.Join(domainFolder, "entities")
		domainUsecasesFolder := filepath.Join(domainFolder, "usecases")

		err = createFolderIfNotExist(domainRepoFolder)
		if err != nil {
			return err
		}

		err = createFolderIfNotExist(domainEntitiesFolder)
		if err != nil {
			return err
		}

		err = createFolderIfNotExist(domainUsecasesFolder)
		if err != nil {
			return err
		}

		domainRepoFile := filepath.Join(domainRepoFolder, feature+"_repo.dart")
		domainRepoContent := fmt.Sprintf(`abstract class %sRepo {

}
`, strings.Title(feature))

		err = createFileWithContent(domainRepoFile, domainRepoContent)
		if err != nil {
			return err
		}

		domainEntityFile := filepath.Join(domainEntitiesFolder, feature+"_entity.dart")
		domainEntityContent := fmt.Sprintf(`class %sEntity {}
`, strings.Title(feature))

		err = createFileWithContent(domainEntityFile, domainEntityContent)
		if err != nil {
			return err
		}

		presentationScreensFolder := filepath.Join(presentationFolder, "screens")
		err = createFolderIfNotExist(presentationScreensFolder)
		if err != nil {
			return err
		}

		presentationScreenFile := filepath.Join(presentationScreensFolder, feature+"_screen.dart")
		presentationScreenContent := fmt.Sprintf(`
		import 'package:flutter/material.dart';

class %sScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("%s")),
      body: Center(child: Text("Welcome to %s screen")),
    );
  }
}
`, strings.Title(feature), strings.Title(feature), strings.Title(feature))

		err = createFileWithContent(presentationScreenFile, presentationScreenContent)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide feature names as arguments (e.g., 'go run main.go auth profile').")
		return
	}

	features := os.Args[1:]

	err := createCoreStructure()
	if err != nil {
		fmt.Println("Error creating core structure:", err)
		return
	}

	err = createFeatureStructure(features)
	if err != nil {
		fmt.Println("Error creating feature structure:", err)
		return
	}

	fmt.Println("Flutter clean architecture template generated successfully!")
}
