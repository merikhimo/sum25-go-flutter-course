import 'dart:async';
import 'dart:convert';
import 'package:http/http.dart' as http;
import '../models/message.dart';

class ApiService {
  // TODO: Add static const String baseUrl = 'http://localhost:8080';
  // TODO: Add static const Duration timeout = Duration(seconds: 30);
  // TODO: Add late http.Client _client field
static const String baseUrl = 'http://localhost:8080';
static const Duration timeout = Duration(seconds: 30);
late http.Client _client;
  // TODO: Add constructor that initializes _client = http.Client();
  ApiService(): _client = http.Client();

  // TODO: Add dispose() method that calls _client.close();
  void dispose() {
    _client.close();
  }

  // TODO: Add _getHeaders() method that returns Map<String, String>
  // Return headers with 'Content-Type': 'application/json' and 'Accept': 'application/json'
  Map<String, String> _getHeaders() {
     return {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  };
    
  }

  // TODO: Add _handleResponse<T>() method with parameters:
  // http.Response response, T Function(Map<String, dynamic>) fromJson
  // Check if response.statusCode is between 200-299
  // If successful, decode JSON and return fromJson(decodedData)
  // If 400-499, throw client error with message from response
  // If 500-599, throw server error
  // For other status codes, throw general error
  T _handleResponse<T>(
  http.Response response,
  T Function(Map<String, dynamic>) fromJson,
  ) {
  final statusCode = response.statusCode;

  if (statusCode >= 200 && statusCode < 300) {
    Map<String, dynamic> data = jsonDecode(response.body);
    return fromJson(data);
  } else if (statusCode >= 400 && statusCode < 500) {
    throw Exception('сlient error: ${statusCode}: ${response.body}');
  } else if (statusCode >= 500 && statusCode < 600) {
    throw Exception('server error: ${statusCode}');
  } else {
    throw Exception('error: ${statusCode}');
  }
}

  // Get all messages
  Future<List<Message>> getMessages() async {
    final URL = Uri.parse('$baseUrl/api/messages');
    try {
    final res = await http.get(URL).timeout(timeout);
    _handleResponse(res, Message.fromJson);
    } on TimeoutException catch (e) {
      throw NetworkException('$e');
    } on http.ClientException catch (e) {
      throw NetworkException('$e');
    }
    // TODO: Implement getMessages
    // Make GET request to '$baseUrl/api/messages'
    // Use _handleResponse to parse response into List<Message>
    // Handle network errors and timeouts
    throw UnimplementedError('TODO: Implement getMessages');
  }

  // Create a new message
  Future<Message> createMessage(CreateMessageRequest request) async {
  request.validate();
  final URL  = Uri.parse('$baseUrl/api/messages');
  final body = jsonEncode(request.toJson());
  try {
    final res = await _client
        .post(URL, headers: _getHeaders(), body: body)
        .timeout(timeout);
    final Map<String, dynamic> wrapper =
        _handleResponse<Map<String, dynamic>>(res, (j) => j);

    return Message.fromJson(wrapper['data'] as Map<String, dynamic>);
  } on TimeoutException {
    throw NetworkException('timed out');
  } on http.ClientException catch (e) {
    throw NetworkException('$e');
  }
  }

  // Update an existing message
  Future<Message> updateMessage(int id, UpdateMessageRequest request) async {
    request.validate();
    final URL = Uri.parse('$baseUrl/api/messages');
    final body = jsonEncode(request.toJson());
    try {
      final res = await _client.put(URL, headers: _getHeaders(), body: body).timeout(timeout);
      final Map<String, dynamic> data = _handleResponse<Map<String, dynamic>>(res, (j) => j);
      return Message.fromJson(data['data'] as Map<String, dynamic>);
    } on TimeoutException {
    throw NetworkException('timed out');
  } on http.ClientException catch (e) {
    throw NetworkException('$e');
  }
    // TODO: Implement updateMessage
    // Validate request using request.validate()
    // Make PUT request to '$baseUrl/api/messages/$id'
    // Include request.toJson() in body
    // Use _handleResponse to parse response
    // Extract message from ApiResponse.data
  }

  // Delete a message
  Future<void> deleteMessage(int id) async {
    // TODO: Implement deleteMessage
    // Make DELETE request to '$baseUrl/api/messages/$id'
    // Check if response.statusCode is 204
    // Throw error if deletion failed
    throw UnimplementedError('TODO: Implement deleteMessage');
  }

  // Get HTTP status information
  Future<HTTPStatusResponse> getHTTPStatus(int statusCode) async {
    // TODO: Implement getHTTPStatus
    // Make GET request to '$baseUrl/api/status/$statusCode'
    // Use _handleResponse to parse response
    // Extract HTTPStatusResponse from ApiResponse.data
    throw UnimplementedError('TODO: Implement getHTTPStatus');
  }

  // Health check
  Future<Map<String, dynamic>> healthCheck() async {
    // TODO: Implement healthCheck
    // Make GET request to '$baseUrl/api/health'
    // Return decoded JSON response
    throw UnimplementedError('TODO: Implement healthCheck');
  }
}

// Custom exceptions
class ApiException implements Exception {
  final String message;
  ApiException(this.message);
  @override
  String toString() {
    // TODO: implement toString
    return 'ApiException: $message';
  }
  // TODO: Add final String message field
  // TODO: Add constructor ApiException(this.message);
  // TODO: Override toString() to return 'ApiException: $message'
}

class NetworkException extends ApiException {
  NetworkException(String message) : super(message);
  // TODO: Add constructor NetworkException(String message) : super(message);
}

class ServerException extends ApiException {
  ServerException(String message) : super(message);
  // TODO: Add constructor ServerException(String message) : super(message);
}

class ValidationException extends ApiException {
  ValidationException(String message) : super(message);
  // TODO: Add constructor ValidationException(String message) : super(message);
}
