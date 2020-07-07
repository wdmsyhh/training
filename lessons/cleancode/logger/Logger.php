<?php

class Logger {
    public static function log($message, $context, $isDebug) {
        $logFile = Yii::$app->getRuntimePath() . DIRECTORY_SEPARATOR . 'logs' . DIRECTORY_SEPARATOR . 'test.log';
        $logPath = dirname($logFile);
        if (!is_dir($logPath)) {
            FileHelper::createDirectory($logPath, 0775, true);
        }

        if (!is_string($message) || !is_array($context)) {
            return false;
        }

        $text = self::formatLog($message, $context, $isDebug);
        self::writeFile($logPath, $logFile, $text);

        chmod($logFile, 0775);

        return true;
    }

    private static function formatLog($message, $context, $isDebug) {
        $time = new DateTime()->format('Y-m-d H:i:s');
        $level = $isDebug ? 'debug' : 'info';
        $text = "[$time] [$level] [$message] $context";
        return $text;
    }

    public static function rotateLog($logPath) {
        $maxIndex = 0;
        $filenames = scandir($logPath);
        foreach ($filenames as $filename) {
            if (is_file($logPath . DIRECTORY_SEPARATOR . $filename && strpos($filename, self::FILE_NAME . '.') !== false)) {
                $newIndex = substr($filename, strlen(self::FILE_NAME . '.'));
                $maxIndex = max($maxIndex, $newIndex);
            }
        }

        while ($maxIndex >= 0) {
            $rotateFile = $logFile . ($maxIndex == 0 ? '' : '.' . $maxIndex);
            @copy($rotateFile, $logFile . '.' . ($maxIndex + 1));
            if ($file = @fopen($rotateFile, 'a')) {
                @ftruncate($file, 0);
                @fclose($file);
            }

            $maxIndex--;
        }
    }

    public static function writeFile($logPath, $logFile, $text) {
        $file = @fopen($logFile, 'a');
        clearstatcache();
        if (filesize($logFile) > 1024 * 1024) {
            $logPath = dirname($logFile);
            
            self::rotateLog($logPath);

            fclose($file);
            file_put_contents($logFile, $text, FILE_APPEND);
            return;
        } 
        fwrite($file, $text);
        fclose($file);
    }
}